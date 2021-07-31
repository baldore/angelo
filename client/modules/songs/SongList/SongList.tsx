import React, { useState } from 'react'
import NextLink from 'next/link'
import {
  Alert,
  AlertIcon,
  Box,
  Button,
  Drawer,
  DrawerBody,
  DrawerContent,
  DrawerHeader,
  DrawerOverlay,
  Heading,
  Input,
  Link,
  Stack,
  useDisclosure,
} from '@chakra-ui/react'
import { FiPlus } from 'react-icons/fi'
import useSWR from 'swr'
import { fetchSongs, insertSong } from 'api/songs'

function NewSongForm({
  isOpen,
  onClose,
  onSuccess,
}: {
  isOpen: boolean
  onClose: () => void
  onSuccess: () => void
}) {
  const [error, setError] = useState('')
  const [name, setName] = useState('')
  const onSetName = (e: React.ChangeEvent<HTMLInputElement>) => {
    setName(e.target.value)
  }

  const onSubmit: React.FormEventHandler = async (e) => {
    e.preventDefault()
    setError('')
    try {
      await insertSong(name)
      onSuccess()
      onClose()
      setName('')
    } catch (err) {
      console.log(err?.response?.data)
      setError(err?.response?.data?.message)
    }
  }

  return (
    <Drawer placement="left" onClose={onClose} isOpen={isOpen}>
      <DrawerOverlay />
      <DrawerContent>
        <DrawerHeader borderBottomWidth="1px" px={4}>
          New Song
        </DrawerHeader>
        <DrawerBody px={4}>
          <form onSubmit={onSubmit}>
            <Stack spacing={3} mt={2}>
              <Input
                placeholder="Name"
                size="sm"
                value={name}
                onChange={onSetName}
              />
              {error && (
                <Alert status="error">
                  <AlertIcon /> {error}
                </Alert>
              )}
              <Button type="submit">Save</Button>
            </Stack>
          </form>
        </DrawerBody>
      </DrawerContent>
    </Drawer>
  )
}

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
          <Box key={id}>
            <NextLink href={`/songs/${id}`} passHref>
              <Link>{name}</Link>
            </NextLink>
          </Box>
        ))}
      </Box>

      <NewSongForm isOpen={isOpen} onClose={onClose} onSuccess={revalidate} />
    </>
  )
}

export default SongList
