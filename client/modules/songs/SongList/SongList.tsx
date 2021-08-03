import React from 'react'
import useSWR from 'swr'
import { Box, Button, Heading, useDisclosure } from '@chakra-ui/react'
import { FiPlus } from 'react-icons/fi'
import { fetchSongs } from 'api/songs'
import NewSongForm from 'modules/songs/NewSongForm/NewSongForm'
import SongRow from '../SongRow/SongRow'

function SongList() {
  const { isOpen, onOpen, onClose } = useDisclosure()
  const { data: songs, error, revalidate } = useSWR('/api/songs', fetchSongs)

  if (error) {
    return <div>Failed to load.</div>
  }

  if (!songs) {
    return <div>!loading...</div>
  }

  return (
    <>
      <Heading mb={2}>Songs</Heading>
      <Button
        colorScheme="blue"
        leftIcon={<FiPlus />}
        mt={2}
        mb={5}
        size="sm"
        onClick={onOpen}
      >
        Create song
      </Button>

      <Box>
        {songs.map((song) => (
          <SongRow key={song.id} song={song} />
        ))}
      </Box>

      <NewSongForm isOpen={isOpen} onClose={onClose} onSuccess={revalidate} />
    </>
  )
}

export default SongList
