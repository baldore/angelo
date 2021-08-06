import React, { useEffect, useRef } from 'react'
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
import { Song } from 'types/song'
import { useImmer } from 'use-immer'
import { patchSongLabels } from 'api/songs'
import { mutate } from 'swr'

type Props = {
  song: Song
}

type Label = {
  name: string
}

function SongRow({ song }: Props) {
  const [state, setState] = useImmer({
    enableInput: false,
    newLabel: '',
  })
  const newLabelInputRef = useRef<HTMLInputElement>(null)

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

  const enableNewLabelInput = () => {
    setState((d) => {
      d.enableInput = true
    })
  }

  const onNewLabelChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setState((d) => {
      d.newLabel = e.target.value
    })
  }

  const onNewLabelBlur = () => {
    setState((d) => {
      d.newLabel = ''
      d.enableInput = false
    })
  }

  const onFormSubmit = async (e: React.FormEvent) => {
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
            <form onSubmit={onFormSubmit}>
              <Input
                ref={newLabelInputRef}
                variant="unstyled"
                value={state.newLabel}
                autoFocus
                placeholder="New label..."
                onChange={onNewLabelChange}
                onBlur={onNewLabelBlur}
              />
            </form>
          </Box>
        ) : (
          <Tag as="button" pl={0} onClick={enableNewLabelInput}>
            <TagLeftIcon as={AiOutlinePlus} />
            <TagLabel>Add label</TagLabel>
          </Tag>
        )}
      </Box>
    </Box>
  )
}

export default SongRow
