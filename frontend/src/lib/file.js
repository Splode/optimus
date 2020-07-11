/**
 * fExt returns the extension of a given file.
 * @param {string} filename - The filename.
 * @returns {string}
 */
export function fExt(filename) {
  return filename.split('.').pop()
}

/**
 * fName returns the name of a given file without its extension.
 * @param {string} filename - The filename.
 * @returns {string}
 */
export function fName(filename) {
  filename = filename.replace(/\\/g, '/')
  return filename.substring(
    filename.lastIndexOf('/') + 1,
    filename.lastIndexOf('.')
  )
}

/**
 * fSize returns a pretty string from a number of bytes.
 * For example, 1024 converts to "1 MB"
 * @param {number} bytes - File size in bytes.
 * @returns {string}
 */
export function fSize(bytes) {
  if (bytes === 0) {
    return '0.00 B'
  }
  const e = Math.floor(Math.log(bytes) / Math.log(1024))
  return (bytes / Math.pow(1024, e)).toFixed(2) + ' ' + ' KMGTP'.charAt(e) + 'B'
}
