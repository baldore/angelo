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

type Props = {
  song: Song
}

type Label = {
  name: string
}

// '/songs/{id}/labels'

function SongRow({ song }: Props) {
  const { id, name, labels } = song
  const [state, setState] = useImmer({
    enableInput: false,
    newLabel: '',
  })

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

  const onFormSubmit = (e: React.FormEvent) => {
    e.preventDefault()

    setState((d) => {
      d.newLabel = ''
      d.enableInput = false
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
        {labels.map(({ name }) => (
          <Tag key={name} colorScheme="cyan" mr={2}>
            <TagLabel>{name}</TagLabel>
            <TagCloseButton />
          </Tag>
        ))}

        {state.enableInput ? (
          <Box display="inline-block">
            <form onSubmit={onFormSubmit}>
              <Input
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
