import axios from 'axios'
import type { Song } from 'types/song'

/**
 * Inserts a new song
 */
export async function insertSong(name: string) {
  const request = await axios.post<Song[]>('/api/songs', {
    name,
  })
  return request.data
}

/**
 * Gets a list of songs
 */
export async function fetchSongs() {
  const request = await axios.get<Song[]>('/api/songs')
  return request.data
}

/**
 * Gets a list of song with a specific id
 * @param {string} id
 */
export async function fetchSongWithId(id: string) {
  const request = await axios.get<Song>(`/api/songs/${id}`)
  return request.data
}
