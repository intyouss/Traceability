import { authAPI } from '@/utils/system/request'

export function getMemoryUsage () {
  return authAPI.get('/system/memory/usage')
}

export function getCpuUsage () {
  return authAPI.get('/system/cpu/usage')
}
