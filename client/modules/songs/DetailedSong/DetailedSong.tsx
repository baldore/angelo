import React, { useEffect, useRef } from 'react'
import useSWR from 'swr'
import { Box, Heading, VStack, Text } from '@chakra-ui/react'

import { fetchSongWithId } from 'api/songs'
import Head from 'next/head'
import * as yt from 'modules/youtube/utils'

const songData = {
  resources: [
    {
      id: 'a18ef1d8-fcdc-43f1-84e0-6613e4189831',
      type: 'youtube',
      title: 'fooo',
      description: 'foo description',
      url: 'https://www.youtube.com/watch?v=ZWYPUgn6yaU',
    },
    {
      id: 'a18ef1d8-fcdc-43f1-84e0-6613e4189832',
      type: 'youtube',
      title: 'fooo',
      description: 'foo description',
      url: 'https://www.youtube.com/watch?v=ZWYPUgn6yaU',
    },
  ],
}

type SongResourceProps = {
  id?: string
  title?: string
  description?: string
  url?: string
}

function SongResource({ id, title, description, url }: SongResourceProps) {
  const player = useRef<yt.Player | null>(null)
  useEffect(() => {
    const ready = async () => {
      await yt.onYoutubeReady()
      player.current = yt.makePlayer({ id: id ?? '', url: url ?? '' })
    }
    ready()
  }, [id, url])

  return (
    <Box key={id}>
      <Heading size="lg">{title}</Heading>
      <Text>{description}</Text>
      <div id={id}></div>
    </Box>
  )
}

type Props = {
  id: number
}

function DetailedSong({ id }: Props) {
  const { data: song, error } = useSWR(['/api/songs', id], (_, id) =>
    fetchSongWithId(id)
  )

  if (error) {
    return <div>Failed to load.</div>
  }

  if (!song) {
    return <div>loading...</div>
  }

  return (
    <>
      <Head>
        <title>{song.name}</title>
      </Head>

      <Heading mb={4}>{song.name}</Heading>
      <Heading size="md">Resources</Heading>
      <VStack spacing={6} alignItems="flex-start">
        {songData?.resources?.map(({ id, title, description, url }) => (
          <SongResource
            key={id}
            id={id}
            title={title}
            description={description}
            url={url}
          />
        ))}
      </VStack>
    </>
  )
}

export default DetailedSong
