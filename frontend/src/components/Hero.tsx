//ref: https://chakra-templates.dev/page-sections/hero
import {
  Stack,
  Flex,
  Button,
  Text,
  VStack,
  useBreakpointValue,
} from "@chakra-ui/react";

import titleImage from "../assets/title.jpeg";
export default function WithBackgroundImage() {
  return (
    <Flex
      w={"full"}
      h={"100vh"}
      backgroundImage={titleImage}
      backgroundSize={"cover"}
      backgroundPosition={"center center"}
    >
      <VStack
        w={"full"}
        justify={"center"}
        px={useBreakpointValue({ base: 4, md: 8 })}
        bgGradient={"linear(to-r, blackAlpha.600, transparent)"}
      >
        <Stack maxW={"2xl"} align={"flex-start"} spacing={6}>
          <Text
            color={"white"}
            fontWeight={700}
            lineHeight={1.2}
            fontSize={useBreakpointValue({ base: "3xl", md: "4xl" })}
          >
            {/* Collateral Inclusive NFT + GamiFi. */}
            {/* Collateral Inclusive NFT + GamiFi. */}
            {/* Having fan to hold actual valued NFTs. */}
            {/* Dear cat lovers, */}
            Now, fans can hold real, money-valued NFTs.
          </Text>
          <Text
            color={"white"}
            fontWeight={700}
            lineHeight={1.2}
            fontSize={useBreakpointValue({ base: "xl", md: "2xl" })}
          >
            {/* It's new challenging protocol which aims to become well-known
            Intelectual Property form the blockchain. */}
            {/* Now, fans can hold real, money-valued NFTs. It’s a brand-new
            protocol that aims to establish itself as a well-known component for
            blockchain intellectual property. */}
            It’s a brand-new protocol that aims to establish itself as a
            well-known component for blockchain intellectual property.
          </Text>
          <Stack direction={"row"}>
            <Button
              bg={"blue.400"}
              rounded={"full"}
              color={"white"}
              _hover={{ bg: "blue.500" }}
            >
              {/* Show me more */}
              Coming soon
            </Button>
            <a href="https://discord.gg/SrWgAZdpDb">
              <Button
                bg={"whiteAlpha.300"}
                rounded={"full"}
                color={"white"}
                _hover={{ bg: "whiteAlpha.500" }}
              >
                Join the discord
              </Button>{" "}
            </a>
          </Stack>
        </Stack>
      </VStack>
    </Flex>
  );
}
