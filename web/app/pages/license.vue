<script setup lang="ts">
import type { TableColumn } from '@nuxt/ui'
import { upperFirst } from 'scule'
import type { Row, Table as TanStackTable } from '@tanstack/table-core'
import type { ApiMessageResponse, ApiResponse } from '~/types/api'
import type { License, LicenseActivateData, LicenseQueryData } from '~/types/license'
import dayjs from 'dayjs'

const UButton = resolveComponent('UButton')
const UBadge = resolveComponent('UBadge')
const UDropdownMenu = resolveComponent('UDropdownMenu')
const UCheckbox = resolveComponent('UCheckbox')

useSeoMeta({
  title: '许可证管理'
})

const toast = useToast()
const table = useTemplateRef<{ tableApi?: TanStackTable<License> }>('table')

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

const { data, status } = await useFetch<ApiResponse<LicenseQueryData>>('/api/license/query', {
  key: 'license-query',
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
const selectedRows = computed<Row<License>[]>(() => table.value?.tableApi?.getFilteredSelectedRowModel().rows || [])
const selectedIds = computed<number[]>(() => selectedRows.value.map((row: Row<License>) => row.original.id))
const selectedCount = computed<number>(() => selectedIds.value.length)
const singleDeleteOpen = ref(false)
const singleDeleteLoading = ref(false)
const singleDeleteLicense = ref<{ id: number, code: string } | null>(null)
const renewOpen = ref(false)
const renewLicense = ref<License | null>(null)

function requestDeleteOneLicense(row: Row<License>) {
  singleDeleteLicense.value = {
    id: row.original.id,
    code: row.original.code
  }
  singleDeleteOpen.value = true
}

async function confirmDeleteOneLicense() {
  if (!singleDeleteLicense.value || singleDeleteLoading.value) return

  try {
    singleDeleteLoading.value = true
    const response = await $fetch<ApiMessageResponse>(`/api/license/delete/${singleDeleteLicense.value.id}`, {
      method: 'POST'
    })

    if (response.code !== 0) {
      throw new Error(response.message || '删除失败')
    }

    toast.add({
      title: '删除成功',
      description: `许可证 ${singleDeleteLicense.value.code} 已删除`,
      color: 'success'
    })
    singleDeleteOpen.value = false
    singleDeleteLicense.value = null
    rowSelection.value = {}
    await refreshNuxtData('license-query')
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

function requestRenewLicense(row: Row<License>) {
  renewLicense.value = row.original
  renewOpen.value = true
}

async function onRenewSuccess() {
  await refreshNuxtData('license-query')
}

async function testActivateLicense(row: Row<License>) {
  try {
    const response = await $fetch<ApiResponse<LicenseActivateData>>('/api/license/activate', {
      method: 'POST',
      body: {
        issuer_id: row.original.issuer_id,
        code: row.original.activation_code.replace(/\s+/g, '')
      }
    })

    if (response.code !== 0) {
      throw new Error(response.message || '激活失败')
    }

    const activationData = response.data
    toast.add({
      title: activationData.ok ? '激活测试成功' : '激活测试失败',
      description: `产品: ${activationData.product_name}, 过期时间: ${dayjs(activationData.expire_at).format('YYYY-MM-DD HH:mm:ss')}`,
      color: activationData.ok ? 'success' : 'warning'
    })
  } catch (error) {
    toast.add({
      title: '激活测试失败',
      description: error instanceof Error ? error.message : '请求失败，请稍后重试',
      color: 'error'
    })
  }
}

function getRowItems(row: Row<License>) {
  return [
    {
      type: 'label',
      label: '操作'
    },
    {
      label: '复制激活码',
      icon: 'i-lucide-copy',
      onSelect() {
        navigator.clipboard.writeText(row.original.activation_code.toString())
        toast.add({
          title: '已复制到剪贴板',
          description: '许可证激活码已复制到剪贴板'
        })
      }
    },
    {
      label: '续期',
      icon: 'i-lucide-refresh-cw',
      onSelect() {
        requestRenewLicense(row)
      }
    },
    {
      label: '测试激活',
      icon: 'i-lucide-play-circle',
      onSelect() {
        testActivateLicense(row)
      }
    },
    {
      type: 'separator'
    },
    {
      label: '删除',
      icon: 'i-lucide-trash',
      color: 'error',
      onSelect() {
        requestDeleteOneLicense(row)
      }
    }
  ]
}

const columns: TableColumn<License>[] = [
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
    accessorKey: 'product_id',
    header: '产品ID'
  },
  {
    accessorKey: 'customer_id',
    header: '客户ID'
  },
  {
    accessorKey: 'issuer_id',
    header: '机构ID'
  },
  {
    accessorKey: 'expire_at',
    header: '过期时间',
    cell: ({ row }) => dayjs(row.getValue('expire_at')).format('YYYY-MM-DD HH:mm:ss')
  },
  {
    accessorKey: 'modules',
    header: '模块'
  },
  {
    accessorKey: 'max_instances',
    header: '最大实例数',
    cell: ({ row }) => row.getValue('max_instances') == 0 ? '不限' : row.getValue('max_instances')
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
    accessorKey: 'remarks',
    header: '备注'
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
  <UDashboardPanel id="license">
    <template #header>
      <UDashboardNavbar title="许可证管理">
        <template #leading>
          <UDashboardSidebarCollapse />
        </template>

        <template #right>
          <LicenseAddModal />
        </template>
      </UDashboardNavbar>
    </template>

    <template #body>
      <div class="flex flex-wrap items-center justify-between gap-1.5">
        <UInput
          v-model="code"
          class="max-w-sm"
          icon="i-lucide-search"
          placeholder="按编码搜索..."
        />

        <div class="flex flex-wrap items-center gap-1.5">
          <LicenseDeleteModal :count="selectedCount" :ids="selectedIds">
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
          </LicenseDeleteModal>

          <USelect
            v-model="statusFilter"
            :items="[
              { label: '全部', value: 'all' },
              { label: '有效', value: 1 },
              { label: '无效', value: 0 }
            ]"
            :ui="{ trailingIcon: 'group-data-[state=open]:rotate-180 transition-transform duration-200' }"
            placeholder="筛选状态"
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
          已选择 {{ table?.tableApi?.getFilteredSelectedRowModel().rows.length || 0 }} 条，共
          {{ total }} 条数据。
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
        :title="`删除许可证 ${singleDeleteLicense?.code || ''}`"
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
              :disabled="singleDeleteLoading || !singleDeleteLicense"
              @click="confirmDeleteOneLicense"
            />
          </div>
        </template>
      </UModal>

      <LicenseRenewModal
        v-model:open="renewOpen"
        :license="renewLicense"
        @success="onRenewSuccess"
      />
    </template>
  </UDashboardPanel>
</template>
