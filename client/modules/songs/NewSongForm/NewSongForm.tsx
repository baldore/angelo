import React, { useCallback, useState } from 'react'
import {
  Alert,
  AlertIcon,
  Button,
  Drawer,
  DrawerBody,
  DrawerContent,
  DrawerHeader,
  DrawerOverlay,
  Input,
  Stack,
} from '@chakra-ui/react'
import { insertSong } from 'api/songs'

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

  const onModalClose = useCallback(() => {
    setName('')
    onClose()
  }, [onClose, setName])

  return (
    <Drawer placement="left" onClose={onModalClose} isOpen={isOpen}>
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

export default NewSongForm
