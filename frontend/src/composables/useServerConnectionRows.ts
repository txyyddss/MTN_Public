import { computed, type Ref } from 'vue'

import type { ConnectionAddress, ConnectionResponse } from '@/types/api'

export interface AddressRow {
  badge: string
  copyValue: string
  label: string
  value: string
}

function isAddressRow(row: AddressRow | null): row is AddressRow {
  return row !== null
}

function buildDirectRow(label: string, badge: string, entry?: ConnectionAddress): AddressRow | null {
  if (!entry) {
    return null
  }

  const host = entry.domain || entry.ip
  const hostWithPort = entry.port && host.endsWith(`:${entry.port}`)
  const value = entry.port && !hostWithPort ? `${host}:${entry.port}` : host

  return {
    label,
    badge,
    value,
    copyValue: value
  }
}

function buildSrvRow(label: string, badge: string, value?: string): AddressRow | null {
  if (!value) {
    return null
  }

  return {
    label,
    badge,
    value,
    copyValue: value
  }
}

export function useServerConnectionRows(
  connection: Readonly<Ref<ConnectionResponse | null>>,
  showAllJava: Readonly<Ref<boolean>>
) {
  const javaRows = computed<AddressRow[]>(() => {
    if (!connection.value) {
      return []
    }

    if (showAllJava.value) {
      return [
        buildDirectRow('Java IPv4', 'IPv4', connection.value.connection?.java_ipv4),
        buildDirectRow('Java IPv6', 'IPv6', connection.value.connection?.java_ipv6)
      ].filter(isAddressRow)
    }

    return [
      buildSrvRow('Java SRV IPv4', 'SRV v4', connection.value.addresses.java_ipv4_srv),
      buildSrvRow('Java SRV IPv6', 'SRV v6', connection.value.addresses.java_ipv6_srv)
    ].filter(isAddressRow)
  })

  const bedrockRows = computed<AddressRow[]>(() =>
    [
      buildDirectRow('Bedrock IPv6', 'IPv6', connection.value?.connection?.bedrock_ipv6),
      buildDirectRow('Bedrock IPv4', 'IPv4', connection.value?.connection?.bedrock_ipv4)
    ].filter(isAddressRow)
  )

  return {
    bedrockRows,
    javaRows
  }
}
