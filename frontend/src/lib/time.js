/**
 * prettyTime returns a human-friendly time representation from milliseconds.
 * @param {number} ms
 * @returns {(string)[]}
 */
export function prettyTime(ms) {
  const seconds = (ms / 1000).toFixed(1)
  const minutes = (ms / (1000 * 60)).toFixed(1)
  const hours = (ms / (1000 * 60 * 60)).toFixed(1)
  const days = (ms / (1000 * 60 * 60 * 24)).toFixed(1)

  if (ms < 1000) {
    return [ms, 'Milliseconds']
  } else if (seconds < 60) {
    return [seconds, 'Seconds']
  } else if (minutes < 60) {
    return [minutes, 'Minutes']
  } else if (hours < 24) {
    return [hours, 'Hours']
  } else {
    return [days, 'Days']
  }
}
