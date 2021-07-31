import Head from 'next/head'
import {
  Alert,
  AlertIcon,
  Button,
  Drawer,
  DrawerBody,
  DrawerContent,
  DrawerHeader,
  DrawerOverlay,
  Heading,
  Input,
  Stack,
  useDisclosure,
} from '@chakra-ui/react'
import { useRouter } from 'next/dist/client/router'
import useSWR from 'swr'
import { fetchSongWithId } from 'api/songs'
import React, { useState } from 'react'

function NewResourceForm({
  isOpen,
  onClose,
  onSuccess,
}: {
  isOpen: boolean
  onClose: () => void
  onSuccess: () => void
}) {
  const [error, setError] = useState('')

  const onSubmit: React.FormEventHandler = async (e) => {
    e.preventDefault()
    setError('')
    try {
      // link text not null,
      // label text not null,
      // description text,
      onSuccess()
      onClose()
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
          New Resource
        </DrawerHeader>
        <DrawerBody px={4}>
          <form onSubmit={onSubmit}>
            <Stack spacing={3} mt={2}>
              <Input placeholder="Name" size="sm" />
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

export default function SongView() {
  const { isOpen, onOpen, onClose } = useDisclosure()
  const { query } = useRouter()
  const { id } = query
  const { data: song, error } = useSWR(['/api/songs', id], (_, id) =>
    fetchSongWithId(id)
  )

  if (error) {
    return <div>Failed to load.</div>
  }

  if (!song) {
    return <div>!loading...</div>
  }

  return (
    <>
      <Head>
        <title>{song.name}</title>
      </Head>
      <Heading mb={4}>{song.name}</Heading>
      <Heading size="md">Resources</Heading>
    </>
  )
}
