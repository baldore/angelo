import Head from 'next/head'
import { Heading } from '@chakra-ui/react'
import { useRouter } from 'next/dist/client/router'

export default function SongView() {
  const { query } = useRouter()
  const { id } = query

  return (
    <>
      <Head>
        <title>Detailed song</title>
      </Head>
      <Heading>Detailed song {id}</Heading>
    </>
  )
}
