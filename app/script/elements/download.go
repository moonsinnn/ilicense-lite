// Package main
// @Description: 从openstreetmap下载空间元素坐标信息数据
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const (
	overpassAPIURL = "http://overpass-api.de/api/interpreter"
	// Overpass QL 查询
	query = `
[out:json];
area[name="北京市"]->.searchArea;
(
  node(area.searchArea)[building];
  way(area.searchArea)[building];   
  relation(area.searchArea)[building];
);
out body;
>;
out skel qt;`
	// 保存数据到指定目录
	jsonFilePath    = "/Users/ibingbo/go/src/gitlab.com/gin-app/data/elements/beijing_osm_data.json"
	geoJsonFilePath = "/Users/ibingbo/go/src/gitlab.com/gin-app/data/elements/beijing_osm_data.geojson"
)

type (
	GeoJsonDataFeatureProperty struct {
		ID            string `json:"@id,omitempty"`
		Type          string `json:"type,omitempty"`
		Building      string `json:"building,omitempty"`
		BuildingLevel string `json:"building:levels,omitempty"`
		Name          string `json:"name,omitempty"`
	}
	GeoJsonDataFeatureGeometry struct {
		Type        string      `json:"type,omitempty"`
		Coordinates interface{} `json:"coordinates,omitempty"`
	}
	GeoJsonDataFeature struct {
		ID         string                      `json:"id,omitempty"`
		Type       string                      `json:"type,omitempty"`
		Properties *GeoJsonDataFeatureProperty `json:"properties,omitempty"`
		Geometry   *GeoJsonDataFeatureGeometry `json:"geometry,omitempty"`
	}
	GeoJsonData struct {
		Type      string                `json:"type,omitempty"`
		Generator string                `json:"generator,omitempty"`
		Copyright string                `json:"copyright,omitempty"`
		Timestamp string                `json:"timestamp,omitempty"`
		Features  []*GeoJsonDataFeature `json:"features,omitempty"`
	}

	JsonDataOSM3S struct {
		TimestampOSMBase   string `json:"timestamp_osm_base"`
		TimestampAreasBase string `json:"timestamp_areas_base"`
		Copyright          string `json:"copyright"`
	}
	JsonDataElementTag struct {
		Building      string `json:"building"`
		BuildingLevel string `json:"building:levels"`
		Name          string `json:"name"`
		Type          string `json:"type"`
	}
	JsonDataElementMember struct {
		Type string `json:"type"`
		Ref  int64  `json:"ref"`
		Role string `json:"role"`
	}
	JsonDataElement struct {
		ID      int64                    `json:"id"`
		Type    string                   `json:"type"`
		Lat     float64                  `json:"lat"`
		Lon     float64                  `json:"lon"`
		Tags    *JsonDataElementTag      `json:"tags"`
		Nodes   []int64                  `json:"nodes"`
		Members []*JsonDataElementMember `json:"members"`
	}
	JsonData struct {
		Version   float64            `json:"version"`
		Generator string             `json:"generator"`
		OSM3S     *JsonDataOSM3S     `json:"osm3s"`
		Elements  []*JsonDataElement `json:"elements"`
	}
)

// 请求数据函数
func fetchDataFromOSM(query string) ([]byte, error) {
	resp, err := http.Post(overpassAPIURL, "application/x-www-form-urlencoded", bytes.NewBufferString(query))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// 保存数据到指定目录的函数
func saveDataToFile(data []byte, filePath string) error {
	// 确保目录存在
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	// 将数据写入文件
	err := ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// 将 Overpass API 数据转换为 GeoJSON
// outer逆序,inner顺序
func convertToGeoJSON(overpassData []byte) ([]byte, error) {
	var jsonData JsonData
	if err := json.Unmarshal(overpassData, &jsonData); err != nil {
		return nil, err
	}

	geoJsonData := &GeoJsonData{
		Type:      "FeatureCollection",
		Generator: jsonData.Generator,
		Copyright: "The data included in this document is from www.openstreetmap.org. The data is made available under ODbL.",
		Timestamp: jsonData.OSM3S.TimestampOSMBase,
		Features:  []*GeoJsonDataFeature{},
	}
	nodeMap := make(map[int64][]float64)
	waysMap := make(map[int64]*JsonDataElement)
	// 找出所有的node
	for _, element := range jsonData.Elements {
		if element.Type == "node" {
			item := &GeoJsonDataFeature{
				ID:   fmt.Sprintf("%s/%d", element.Type, element.ID),
				Type: "Feature",
				Properties: &GeoJsonDataFeatureProperty{
					ID: fmt.Sprintf("%s/%d", element.Type, element.ID),
				},
				Geometry: &GeoJsonDataFeatureGeometry{
					Type:        "Point",
					Coordinates: []float64{element.Lon, element.Lat},
				},
			}
			if element.Tags != nil {
				item.Properties.Type = element.Tags.Type
				item.Properties.Name = element.Tags.Name
				item.Properties.Building = element.Tags.Building
				item.Properties.BuildingLevel = element.Tags.BuildingLevel
			}
			geoJsonData.Features = append(geoJsonData.Features, item)
			nodeMap[element.ID] = []float64{element.Lon, element.Lat}
		}
	}
	// 找出所有的way
	for _, element := range jsonData.Elements {
		if element.Type == "way" {
			item := &GeoJsonDataFeature{
				ID:   fmt.Sprintf("%s/%d", element.Type, element.ID),
				Type: "Feature",
				Properties: &GeoJsonDataFeatureProperty{
					ID: fmt.Sprintf("%s/%d", element.Type, element.ID),
				},
				Geometry: &GeoJsonDataFeatureGeometry{
					Type: "LineString",
				},
			}
			if element.Tags != nil {
				item.Properties.Type = element.Tags.Type
				item.Properties.Name = element.Tags.Name
				item.Properties.Building = element.Tags.Building
				item.Properties.BuildingLevel = element.Tags.BuildingLevel
			}
			if len(element.Nodes) > 0 {
				var coordinates [][]float64
				for _, nodeID := range element.Nodes {
					if coord, ok := nodeMap[nodeID]; ok {
						coordinates = append(coordinates, coord)
					}
				}
				// 如果首尾相同，则为 Polygon，否则为 LineString
				if len(element.Nodes) > 2 && element.Nodes[0] == element.Nodes[len(element.Nodes)-1] {
					item.Geometry.Type = "Polygon"
					item.Geometry.Coordinates = [][][]float64{coordinates} // Polygon 需要多一层数组嵌套
				} else {
					item.Geometry.Coordinates = coordinates
				}
			}
			waysMap[element.ID] = element
			geoJsonData.Features = append(geoJsonData.Features, item)
		}
	}

	// 找出所有的relation
	for _, element := range jsonData.Elements {
		if element.Type == "relation" {
			item := &GeoJsonDataFeature{
				ID:   fmt.Sprintf("%s/%d", element.Type, element.ID),
				Type: "Feature",
				Properties: &GeoJsonDataFeatureProperty{
					ID: fmt.Sprintf("%s/%d", element.Type, element.ID),
				},
				Geometry: &GeoJsonDataFeatureGeometry{
					Type: "Polygon",
				},
			}
			if element.Tags != nil {
				item.Properties.Type = element.Tags.Type
				item.Properties.Name = element.Tags.Name
				item.Properties.Building = element.Tags.Building
				item.Properties.BuildingLevel = element.Tags.BuildingLevel
			}
			if len(element.Members) > 0 {
				var outer [][]float64
				var inners [][][]float64
				for _, member := range element.Members {
					if way, ok := waysMap[member.Ref]; ok {
						if member.Role == "outer" {
							for i := len(way.Nodes) - 1; i >= 0; i-- {
								if coord, ok := nodeMap[way.Nodes[i]]; ok {
									outer = append(outer, coord)
								}
							}
						} else {
							var cList [][]float64
							for _, nodeID := range way.Nodes {
								if coord, ok := nodeMap[nodeID]; ok {
									cList = append(cList, coord)
								}
							}
							inners = append(inners, cList)
						}
					}
				}
				item.Geometry.Coordinates = append([][][]float64{outer}, inners...)
			}
			geoJsonData.Features = append(geoJsonData.Features, item)
		}
	}
	data, err := json.Marshal(geoJsonData)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func main() {

	log.Println("downloading geo json data...")
	// 获取数据
	data, err := fetchDataFromOSM(query)
	if err != nil {
		log.Println("fetching data from osm error:", err)
		return
	}

	if err := saveDataToFile(data, jsonFilePath); err != nil {
		log.Println("saving json data to file:", err)
		return
	}
	log.Printf("json data successfully saved to %s\n", jsonFilePath)

	geoJsonData, err := convertToGeoJSON(data)
	if err != nil {
		log.Println("converting to geojson error:", err)
		return
	}

	if err := saveDataToFile(geoJsonData, geoJsonFilePath); err != nil {
		log.Println("saving json data to file:", err)
		return
	}
	log.Printf("geo json data successfully saved to %s\n", geoJsonFilePath)

}
