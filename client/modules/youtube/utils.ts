declare global {
  interface Window {
    YT: YoutubeAPI
  }
}

type YoutubeAPI = {
  Player: Player
}

export type Player = {
  new (id: string, options: PlayerOptions): Player
}

type PlayerOptions = {
  videoId: string
  width?: string
  height?: string
}

/**
 * Checks if youtube iframe API is ready to be used
 */
export function onYoutubeReady(): Promise<void> {
  return new Promise((resolve, _reject) => {
    if (process.browser) {
      const w = window as any
      if (w.YT) {
        resolve()
      } else {
        w.onYoutubeReady = function () {
          resolve()
        }
      }
    }
  })
}

export function getId(url: string): string {
  const regExp =
    /^.*((youtu.be\/)|(v\/)|(\/u\/\w\/)|(embed\/)|(watch\?))\??v?=?([^#&?]*).*/
  const match = url.match(regExp)

  return match && match[7].length == 11 ? match[7] : ''
}

type MakePlayerOptions = {
  id: string
  url: string
}

export function makePlayer({ id, url }: MakePlayerOptions): Player {
  return new window.YT.Player(id, {
    videoId: getId(url),
    width: '300',
    height: '200',
  })
}
