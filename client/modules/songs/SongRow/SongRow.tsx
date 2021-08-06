import React, { useRef } from 'react'
import NextLink from 'next/link'
import {
  Alert,
  AlertDescription,
  AlertIcon,
  AlertTitle,
  Box,
  CloseButton,
  Input,
  Link,
  Tag,
  TagCloseButton,
  TagLabel,
  TagLeftIcon,
} from '@chakra-ui/react'
import { AiOutlinePlus } from 'react-icons/ai'
import { Label, Song } from 'types/song'
import { useImmer } from 'use-immer'
import { patchSongLabels } from 'api/songs'
import { mutate } from 'swr'

function SongRow({ song }: { song: Song }) {
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

    if (labelExists) return

    // TODO: Add service to show popup errors/alerts

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
      <Alert status="success" variant="left-accent">
        <AlertIcon />
        <AlertTitle mr={2}>Your browser is outdated!</AlertTitle>
        <AlertDescription>
          Your Chakra experience may be degraded.
        </AlertDescription>
        <CloseButton position="absolute" right="8px" top="8px" />
      </Alert>
    </Box>
  )
}

export default SongRow
