// src/composables/useFeedBus.js
const listeners = new Set()

export function usePostShareBus() {
  function onPostShared(callback) {
    listeners.add(callback)
    return () => listeners.delete(callback)
  }

  function emitPostShared(repostedPost) {
    listeners.forEach((cb) => {
      try {
        cb(repostedPost)
      } catch (e) {
        console.error('usePostShareBus listener failed', e)
      }
    })
  }

  return { onPostShared, emitPostShared }
}