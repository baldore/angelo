import NextLink from 'next/link'
import { Link } from '@chakra-ui/react'
import { useRouter } from 'next/dist/client/router'
import React from 'react'
import DetailedSong from 'modules/songs/DetailedSong/DetailedSong'

export default function SongView() {
  const router = useRouter()

  if (!router.isReady) {
    return null
  }

  const { id } = router.query

  return (
    <>
      <NextLink href="/" passHref>
        <Link>{'<'} Go back to songs</Link>
      </NextLink>

      <DetailedSong id={Number(id)} />
    </>
  )
}
