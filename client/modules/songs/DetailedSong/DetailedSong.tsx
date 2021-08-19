import React from 'react'
import useSWR from 'swr'
import { Box, Heading, VStack, Text } from '@chakra-ui/react'

import { fetchSongWithId } from 'api/songs'
import Head from 'next/head'

const songData = {
  resources: [
    {
      id: 1,
      type: 'youtube',
      title: 'fooo',
      description: 'foo description',
      url: 'https://www.youtube.com/watch?v=ZWYPUgn6yaU',
    },
    {
      id: 2,
      type: 'youtube',
      title: 'fooo',
      description: 'foo description',
      url: 'https://www.youtube.com/watch?v=ZWYPUgn6yaU',
    },
  ],
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
      <VStack spacing={2} alignItems="flex-start">
        {songData?.resources?.map((resource) => (
          <Box key={resource.id}>
            <Heading>{resource?.title}</Heading>
            <Text>{resource?.description}</Text>
          </Box>
        ))}
      </VStack>
    </>
  )
}

export default DetailedSong
