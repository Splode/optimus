/**
 * fExt returns the extension of a given file.
 * @param {string} filename
 * @returns {string}
 */
export function fExt(filename) {
    return filename.split('.').pop()
}

/**
 * fName returns the name of a given file without its extension.
 * @param {string} filename
 * @returns {string}
 */
export function fName(filename) {
    filename = filename.replace(/\\/g, '/');
    return filename.substring(filename.lastIndexOf('/') + 1, filename.lastIndexOf('.'));
}
