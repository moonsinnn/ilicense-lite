<script setup lang="ts">
import type { TableColumn } from '@nuxt/ui'
import { upperFirst } from 'scule'
import type { Row, Table as TanStackTable } from '@tanstack/table-core'
import type { ApiMessageResponse, ApiResponse } from '~/types/api'
import type { Product, ProductQueryData } from '~/types/product'
import dayjs from 'dayjs'

const UButton = resolveComponent('UButton')
const UBadge = resolveComponent('UBadge')
const UDropdownMenu = resolveComponent('UDropdownMenu')
const UCheckbox = resolveComponent('UCheckbox')

const toast = useToast()
const table = useTemplateRef<{ tableApi?: TanStackTable<Product> }>('table')

const columnFilters = ref([{
  id: 'code',
  value: ''
}])
const columnVisibility = ref()
const rowSelection = ref<Record<string, boolean>>({})

const pagination = ref({
  pageIndex: 0,
  pageSize: 10
})

const requestBody = computed(() => ({
  page: pagination.value.pageIndex + 1,
  size: pagination.value.pageSize
}))

const { data, status } = await useFetch<ApiResponse<ProductQueryData>>('/api/product/query', {
  key: 'product-query',
  method: 'POST',
  body: requestBody,
  watch: [requestBody],
  default: () => ({
    code: 0,
    message: 'ok',
    data: {
      total: 0,
      items: []
    }
  }),
  lazy: true
})

const tableData = computed(() => data.value?.data?.items || [])
const total = computed(() => data.value?.data?.total || 0)
const selectedRows = computed<Row<Product>[]>(() => table.value?.tableApi?.getFilteredSelectedRowModel().rows || [])
const selectedIds = computed<number[]>(() => selectedRows.value.map((row: Row<Product>) => row.original.id))
const selectedCount = computed<number>(() => selectedIds.value.length)
const singleDeleteOpen = ref(false)
const singleDeleteLoading = ref(false)
const singleDeleteProduct = ref<{ id: number, name: string } | null>(null)

function requestDeleteOneProduct(row: Row<Product>) {
  singleDeleteProduct.value = {
    id: row.original.id,
    name: row.original.name
  }
  singleDeleteOpen.value = true
}

async function confirmDeleteOneProduct() {
  if (!singleDeleteProduct.value || singleDeleteLoading.value) return

  try {
    singleDeleteLoading.value = true
    const response = await $fetch<ApiMessageResponse>(`/api/product/delete/${singleDeleteProduct.value.id}`, {
      method: 'POST'
    })

    if (response.code !== 0) {
      throw new Error(response.message || '删除失败')
    }

    toast.add({
      title: '删除成功',
      description: `机构 ${singleDeleteProduct.value.name} 已删除`,
      color: 'success'
    })
    singleDeleteOpen.value = false
    singleDeleteProduct.value = null
    rowSelection.value = {}
    await refreshNuxtData('product-query')
  } catch (error) {
    toast.add({
      title: '删除失败',
      description: error instanceof Error ? error.message : '请求失败，请稍后重试',
      color: 'error'
    })
  } finally {
    singleDeleteLoading.value = false
  }
}

function getRowItems(row: Row<Product>) {
  return [
    {
      type: 'label',
      label: '操作'
    },
    {
      type: 'separator'
    },
    {
      label: '删除',
      icon: 'i-lucide-trash',
      color: 'error',
      onSelect() {
        requestDeleteOneProduct(row)
      }
    }
  ]
}

const columns: TableColumn<Product>[] = [
  {
    id: 'select',
    header: ({ table }) =>
      h(UCheckbox, {
        'modelValue': table.getIsSomePageRowsSelected()
          ? 'indeterminate'
          : table.getIsAllPageRowsSelected(),
        'onUpdate:modelValue': (value: boolean | 'indeterminate') =>
          table.toggleAllPageRowsSelected(!!value),
        'ariaLabel': 'Select all'
      }),
    cell: ({ row }) =>
      h(UCheckbox, {
        'modelValue': row.getIsSelected(),
        'onUpdate:modelValue': (value: boolean | 'indeterminate') => row.toggleSelected(!!value),
        'ariaLabel': 'Select row'
      })
  },
  {
    accessorKey: 'id',
    header: 'ID'
  },
  {
    accessorKey: 'code',
    header: '编码'
  },
  {
    accessorKey: 'name',
    header: '名称'
  },
  {
    accessorKey: 'description',
    header: '描述'
  },
  {
    accessorKey: 'status',
    header: '状态',
    filterFn: 'equals',
    cell: ({ row }) => {
      const color = {
        1: 'success' as const,
        0: 'neutral' as const
      }[row.original.status]

      return h(UBadge, { class: 'capitalize', variant: 'subtle', color }, () =>
        row.original.status == 1 ? '有效' : '无效'
      )
    }
  },
  {
    accessorKey: 'created_at',
    header: '创建时间',
    cell: ({ row }) => {
      return dayjs(row.getValue('created_at')).format('YYYY-MM-DD HH:mm:ss')
    }
  },
  {
    accessorKey: 'updated_at',
    header: '更新时间',
    cell: ({ row }) => {
      return dayjs(row.getValue('updated_at')).format('YYYY-MM-DD HH:mm:ss')
    }
  },
  {
    id: 'actions',
    cell: ({ row }) => {
      return h(
        'div',
        { class: 'text-right' },
        h(
          UDropdownMenu,
          {
            content: {
              align: 'end'
            },
            items: getRowItems(row)
          },
          () =>
            h(UButton, {
              icon: 'i-lucide-ellipsis-vertical',
              color: 'neutral',
              variant: 'ghost',
              class: 'ml-auto'
            })
        )
      )
    }
  }
]

const statusFilter = ref('all')

watch(() => statusFilter.value, (newVal) => {
  if (!table?.value?.tableApi) return

  const statusColumn = table.value.tableApi.getColumn('status')
  if (!statusColumn) return

  if (newVal === 'all') {
    statusColumn.setFilterValue(undefined)
  } else {
    statusColumn.setFilterValue(newVal)
  }
})

const code = computed({
  get: (): string => {
    return (table.value?.tableApi?.getColumn('code')?.getFilterValue() as string) || ''
  },
  set: (value: string) => {
    table.value?.tableApi?.getColumn('code')?.setFilterValue(value || undefined)
  }
})
</script>

<template>
  <UDashboardPanel id="product">
    <template #header>
      <UDashboardNavbar title="Product">
        <template #leading>
          <UDashboardSidebarCollapse />
        </template>

        <template #right>
          <ProductAddModal />
        </template>
      </UDashboardNavbar>
    </template>

    <template #body>
      <div class="flex flex-wrap items-center justify-between gap-1.5">
        <UInput
          v-model="code"
          class="max-w-sm"
          icon="i-lucide-search"
          placeholder="Filter code..."
        />

        <div class="flex flex-wrap items-center gap-1.5">
          <ProductDeleteModal :count="selectedCount" :ids="selectedIds">
            <UButton
              v-if="selectedCount"
              label="删除"
              color="error"
              variant="subtle"
              icon="i-lucide-trash"
            >
              <template #trailing>
                <UKbd>
                  {{ selectedCount }}
                </UKbd>
              </template>
            </UButton>
          </ProductDeleteModal>

          <USelect
            v-model="statusFilter"
            :items="[
              { label: '全部', value: 'all' },
              { label: '有效', value: 1 },
              { label: '无效', value: 0 }
            ]"
            :ui="{ trailingIcon: 'group-data-[state=open]:rotate-180 transition-transform duration-200' }"
            placeholder="Filter status"
            class="min-w-28"
          />
          <UDropdownMenu
            :items="
              table?.tableApi
                ?.getAllColumns()
                .filter((column: any) => column.getCanHide())
                .map((column: any) => ({
                  label: upperFirst(column.id),
                  type: 'checkbox' as const,
                  checked: column.getIsVisible(),
                  onUpdateChecked(checked: boolean) {
                    table?.tableApi?.getColumn(column.id)?.toggleVisibility(!!checked)
                  },
                  onSelect(e?: Event) {
                    e?.preventDefault()
                  }
                }))
            "
            :content="{ align: 'end' }"
          >
            <UButton
              label="展示"
              color="neutral"
              variant="outline"
              trailing-icon="i-lucide-settings-2"
            />
          </UDropdownMenu>
        </div>
      </div>

      <UTable
        ref="table"
        v-model:column-filters="columnFilters"
        v-model:column-visibility="columnVisibility"
        v-model:row-selection="rowSelection"
        class="shrink-0"
        :data="tableData"
        :columns="columns"
        :loading="status === 'pending'"
        :ui="{
          base: 'table-fixed border-separate border-spacing-0',
          thead: '[&>tr]:bg-elevated/50 [&>tr]:after:content-none',
          tbody: '[&>tr]:last:[&>td]:border-b-0',
          th: 'py-2 first:rounded-l-lg last:rounded-r-lg border-y border-default first:border-l last:border-r',
          td: 'border-b border-default',
          separator: 'h-0'
        }"
      />

      <div class="flex items-center justify-between gap-3 border-t border-default pt-4 mt-auto">
        <div class="text-sm text-muted">
          {{ table?.tableApi?.getFilteredSelectedRowModel().rows.length || 0 }} of
          {{ total }} row(s) total.
        </div>

        <div class="flex items-center gap-1.5">
          <UPagination
            :page="pagination.pageIndex + 1"
            :items-per-page="pagination.pageSize"
            :total="total"
            @update:page="(p: number) => (pagination.pageIndex = p - 1)"
          />
        </div>
      </div>

      <UModal
        v-model:open="singleDeleteOpen"
        :title="`删除机构 ${singleDeleteProduct?.name || ''}`"
        description="你确定吗, 该项目操作不可恢复."
      >
        <template #body>
          <div class="flex justify-end gap-2">
            <UButton
              label="取消"
              color="neutral"
              variant="subtle"
              @click="singleDeleteOpen = false"
            />
            <UButton
              label="删除"
              color="error"
              variant="solid"
              :loading="singleDeleteLoading"
              :disabled="singleDeleteLoading || !singleDeleteProduct"
              @click="confirmDeleteOneProduct"
            />
          </div>
        </template>
      </UModal>
    </template>
  </UDashboardPanel>
</template>
