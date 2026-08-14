// src/composables/useFeedBus.js
const listeners = new Set()

export function usePostShareBus() {
  function onPostShared(callback) {
    listeners.add(callback)
    // returns an unsubscribe function — call this in onUnmounted()
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