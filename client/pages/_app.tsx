import React from 'react'
import type { AppProps } from 'next/app'
import { ChakraProvider, Container, VStack } from '@chakra-ui/react'

import Notifications from 'modules/notifications/Notifications'

/**
 * Used to set the layout of the page.
 */
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
      <Notifications />
    </VStack>
  )
}

function App({ Component, pageProps }: AppProps) {
  return (
    <ChakraProvider>
      <MainContainer>
        <Component {...pageProps} />
      </MainContainer>
    </ChakraProvider>
  )
}

export default App
