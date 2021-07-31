import React from 'react'
import NextLink from 'next/link'
import {
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
import { songs } from 'data/songs'
import { FiPlus } from 'react-icons/fi'

function NewSongForm({
  isOpen,
  onClose,
}: {
  isOpen: boolean
  onClose: () => void
}) {
  const onSubmit: React.FormEventHandler = (e) => {
    e.preventDefault()

    // TODO: Add logic to create new song
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
              <Input placeholder="Name" size="sm" />
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

      <NewSongForm isOpen={isOpen} onClose={onClose} />
    </>
  )
}

export default SongList
