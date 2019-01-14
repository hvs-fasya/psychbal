import { Notify } from 'quasar'

export function notify401 () {
  Notify.create('Ошибка авторизации')
}

export function notify500 () {
  Notify.create('Что-то пошло не так')
}

export function capitalize(s){
  return s.charAt(0).toUpperCase() + s.slice(1)
}
