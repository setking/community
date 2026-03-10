type messageType = {
  text: string
  color: string
}
export const useMessagesStore = defineStore('messages', () => {
  const queue = ref<any[]>([])
  function add (message: messageType) {
    queue.value.push(message)
  }

  return { queue, add }
})
