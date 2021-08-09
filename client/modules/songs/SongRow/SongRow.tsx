import { mutate } from 'swr'
import { useAtom } from 'jotai'
import React, { useRef } from 'react'
import NextLink from 'next/link'
import {
  Box,
  Input,
  Link,
  Tag,
  TagCloseButton,
  TagLabel,
  TagLeftIcon,
} from '@chakra-ui/react'
import { AiOutlinePlus } from 'react-icons/ai'
import { useImmer } from 'use-immer'

import { Label, Song } from 'types/song'
import { patchSongLabels } from 'api/songs'
import { addNotificationAtom } from 'modules/notifications/Notifications'

function SongRow({ song }: { song: Song }) {
  const [, addNotification] = useAtom(addNotificationAtom)
  const newLabelInputRef = useRef<HTMLInputElement>(null)
  const [state, setState] = useImmer({
    enableInput: false,
    newLabel: '',
  })

  const { id, name, labels } = song

  const updateLabels = async (newLabels: Label[]) => {
    try {
      await patchSongLabels(id, newLabels)
      mutate('/api/songs')
    } catch (err) {
      console.error(err)
    }
  }

  const deleteLabel = async ({ name }: Label) => {
    const newLabels = labels.filter((label) => label.name != name)
    await updateLabels(newLabels)
  }

  const enableNewLabelForm = () => {
    setState((d) => {
      d.enableInput = true
    })
  }

  const updateNewLabelValue = (e: React.ChangeEvent<HTMLInputElement>) => {
    setState((d) => {
      d.newLabel = e.target.value
    })
  }

  const resetForm = () => {
    setState((d) => {
      d.newLabel = ''
      d.enableInput = false
    })
  }

  const createNewLabel = async (e: React.FormEvent) => {
    e.preventDefault()

    const { newLabel } = state
    const labelExists = Boolean(labels.find((l) => l.name === newLabel))

    if (labelExists) {
      addNotification({
        message: 'Label already exists...',
      })
      return
    }

    const newLabels = labels.concat({ name: newLabel })
    await updateLabels(newLabels)

    setState((d) => {
      d.newLabel = ''
    })
  }

  return (
    <Box key={id} mb={3}>
      <NextLink href={`/songs/${id}`} passHref>
        <Link fontSize="xl" textTransform="capitalize">
          {id}: {name}
        </Link>
      </NextLink>
      <Box>
        {labels.map((label) => (
          <Tag key={label.name} colorScheme="cyan" mr={2}>
            <TagLabel>{label.name}</TagLabel>
            <TagCloseButton onClick={() => deleteLabel(label)} />
          </Tag>
        ))}

        {state.enableInput ? (
          <Box display="inline-block">
            <form onSubmit={createNewLabel}>
              <Input
                ref={newLabelInputRef}
                variant="unstyled"
                value={state.newLabel}
                autoFocus
                placeholder="New label..."
                onChange={updateNewLabelValue}
                onBlur={resetForm}
              />
            </form>
          </Box>
        ) : (
          <Tag as="button" pl={0} onClick={enableNewLabelForm}>
            <TagLeftIcon as={AiOutlinePlus} />
            <TagLabel>Add label</TagLabel>
          </Tag>
        )}
      </Box>
    </Box>
  )
}

export default SongRow
