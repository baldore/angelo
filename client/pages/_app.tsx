import '../styles/globals.css'
import { ChakraProvider, Container, VStack } from '@chakra-ui/react'
import type { AppProps } from 'next/app'
import React from 'react'

function MainContainer({ children }: { children: React.ReactElement }) {
  return (
    <VStack h="100vh" py={4}>
      <Container
        backgroundColor="gray.100"
        maxW="container.lg"
        minH="100%"
        p={7}
      >
        {children}
      </Container>
    </VStack>
  )
}

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <ChakraProvider>
      <MainContainer>
        <Component {...pageProps} />
      </MainContainer>
    </ChakraProvider>
  )
}

export default MyApp
