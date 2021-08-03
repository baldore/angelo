import React from 'react'
import NextLink from 'next/link'
import useSWR from 'swr'
import {
  Box,
  Button,
  Heading,
  Link,
  Tag,
  TagLabel,
  TagLeftIcon,
  useDisclosure,
} from '@chakra-ui/react'
import { FiPlus } from 'react-icons/fi'
import { AiOutlinePlus } from 'react-icons/ai'
import { fetchSongs } from 'api/songs'
import NewSongForm from 'modules/songs/NewSongForm/NewSongForm'

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
        {songs.map(({ id, name }) => (
          <Box key={id} mb={3}>
            <NextLink href={`/songs/${id}`} passHref>
              <Link fontSize="xl" textTransform="capitalize">
                {name}
              </Link>
            </NextLink>
            <Box>
              <Tag>
                <TagLeftIcon as={AiOutlinePlus} />
                <TagLabel>Add label</TagLabel>
              </Tag>
            </Box>
          </Box>
        ))}
      </Box>

      <NewSongForm isOpen={isOpen} onClose={onClose} onSuccess={revalidate} />
    </>
  )
}

export default SongList
