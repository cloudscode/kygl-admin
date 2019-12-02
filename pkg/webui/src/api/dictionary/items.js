import request from '@/utils/request'
import qs from 'qs'

export function fetchDictionaryItemsList(query) {
  return request({
    url: 'v1/dictionary/items',
    method: 'get',
    params: query
  })
}

export function fetchDictionaryItems(id) {
  return request({
    url: 'v1/dictionary/items',
    method: 'get',
    params: {
      id
    }
  })
}

export function createDictionaryItems(data) {
  return request({
    url: 'v1/dictionary/items',
    method: 'post',
    data: qs.stringify(data)
  })
}

export function updateDictionaryItems(data) {
  return request({
    url: 'v1/dictionary/items/' + data.id,
    method: 'put',
    data: qs.stringify(data)
  })
}

export function deleteDictionaryItems(data) {
  return request({
    url: 'v1/dictionary/items/' + data.id,
    method: 'delete',
    data: qs.stringify(data)
  })
}