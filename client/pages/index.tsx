import React from 'react'
import Head from 'next/head'
import SongList from 'modules/songs/SongList/SongList'

export default function Home() {
  return (
    <>
      <Head>
        <title>Songs</title>
      </Head>
      <SongList />
    </>
  )
}
