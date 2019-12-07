import request from '@/utils/request'
import qs from 'qs'

export function fetchDisquisionList(query) {
  return request({
    url: 'v1/disquisions',
    method: 'get',
    params: query
  })
}

export function fetchDisquision(id) {
  return request({
    url: 'v1/disquisions',
    method: 'get',
    params: {
      id
    }
  })
}

export function createDisquision(data) {
  return request({
    url: 'v1/disquisions',
    method: 'post',
    data: qs.stringify(data)
  })
}

export function updateDisquision(data) {
  return request({
    url: 'v1/disquisions/' + data.id,
    method: 'put',
    data: qs.stringify(data)
  })
}

export function deleteDisquision(data) {
  return request({
    url: 'v1/disquisions/' + data.id,
    method: 'delete',
    data: qs.stringify(data)
  })
}

