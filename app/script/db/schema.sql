-- 产品表
CREATE TABLE IF NOT EXISTS product
(
    id          BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    code        VARCHAR(50) UNIQUE NOT NULL DEFAULT '' COMMENT '产品代码, 唯一标识',
    name        VARCHAR(100)       NOT NULL DEFAULT '' COMMENT '产品名称',
    description TEXT COMMENT '产品描述',
    status      TINYINT UNSIGNED   NOT NULL DEFAULT 1 COMMENT '状态: 1启用, 0禁用',
    created_at  DATETIME           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at  DATETIME           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='产品表';

-- 客户表
CREATE TABLE IF NOT EXISTS customer
(
    id             BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    code           VARCHAR(50) UNIQUE NOT NULL DEFAULT '' COMMENT '客户编号',
    name           VARCHAR(100)       NOT NULL DEFAULT '' COMMENT '客户名称',
    contact_person VARCHAR(50)                 DEFAULT '' COMMENT '联系人',
    phone          VARCHAR(20)                 DEFAULT '' COMMENT '联系电话',
    email          VARCHAR(100)                DEFAULT '' COMMENT '联系邮箱',
    address        TEXT COMMENT '公司地址',
    status         TINYINT UNSIGNED   NOT NULL DEFAULT 1 COMMENT '状态: 1正常, 0禁用',
    created_at     DATETIME           NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at     DATETIME           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='客户表';

-- License记录表
CREATE TABLE IF NOT EXISTS license
(
    id              BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    code            VARCHAR(50) UNIQUE NOT NULL DEFAULT '' COMMENT 'License编号, 系统唯一',
    product_id      BIGINT UNSIGNED    NOT NULL COMMENT '产品ID, 外键关联 product',
    customer_id     BIGINT UNSIGNED    NOT NULL COMMENT '客户ID, 外键关联 customer',
    activation_code TEXT               NOT NULL COMMENT '激活码(Base64 URL安全编码)',
    issue_date      DATE               NOT NULL COMMENT '签发日期',
    expiry_date     DATE               NOT NULL COMMENT '到期日期',
    modules         TEXT COMMENT '授权模块(JSON)',
    max_instances   INT                NOT NULL DEFAULT -1 COMMENT '最大实例数, -1 表示无限',
    status          TINYINT UNSIGNED   NOT NULL DEFAULT 1 COMMENT '状态: 1有效, 0无效',
    remarks         TEXT COMMENT '备注',
    created_at      DATETIME           NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      DATETIME           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_customer_product (customer_id, product_id),
    INDEX idx_expiry_date (expiry_date)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='License记录表';



