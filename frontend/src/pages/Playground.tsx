//ref: https://chakra-templates.dev/page-sections/hero
import {
  Stack,
  Flex,
  Button,
  Text,
  VStack,
  useBreakpointValue,
  Spacer,
  Select,
  FormControl,
  FormLabel,
  FormControlOptions,
  Grid,
  GridItem,
  Heading,
} from "@chakra-ui/react";
import React, { useState } from "react";

export default function Playground() {
  const generateTwinkleCats = () => {
    console.log("generateTwinkleCats");
  };
  const [count, setCount] = useState(0);

  return (
    <>
      <Stack spacing={3}>
        <Heading as="h3" size="xl">
          Playground
        </Heading>
        <Heading as="h4" size="l">
          Explore
        </Heading>
        <Text>
          The playground was built using the Twinkle Protocol. Twinkle's traits
          are determined by the Twinkle Seed and additional equipments. The seed
          was generated using twinkle-assets and rendered using the twinkle-sdk.
        </Text>
        {/* {list.forEach((a) => {
          return <>aa{a}</>;
        })} */}
        <Grid
          templateRows="repeat(2, 1fr)"
          templateColumns="repeat(5, 1fr)"
          gap={4}
        >
          <GridItem rowSpan={20} colSpan={1} bg="">
            <Button
              display={{ base: "inline-flex", md: "inline-flex" }}
              fontSize={"sm"}
              fontWeight={600}
              color={"white"}
              bg={"pink.500"}
              _hover={{
                bg: "pink.400",
              }}
              size="lg"
              onClick={generateTwinkleCats}
            >
              <Spacer width={2} />
              Generate Twinkle Cats
            </Button>
          </GridItem>
          <GridItem colSpan={1} bg="">
            <img src="https://imgapi.twinkle.cat/v0/img/randomcat?13" />
          </GridItem>
          <GridItem colSpan={1} bg="">
            <img src="https://imgapi.twinkle.cat/v0/img/randomcat?12" />
          </GridItem>
          <GridItem colSpan={1} bg="">
            <img src="https://imgapi.twinkle.cat/v0/img/randomcat?11" />
          </GridItem>
          <GridItem colSpan={1} bg="">
            <img src="https://imgapi.twinkle.cat/v0/img/randomcat?2" />
          </GridItem>
          <GridItem colSpan={1} bg="">
            <img src="https://imgapi.twinkle.cat/v0/img/randomcat?3" />
          </GridItem>
          <GridItem colSpan={1} bg="">
            <img src="https://imgapi.twinkle.cat/v0/img/randomcat?4" />
          </GridItem>
          <GridItem colSpan={1} bg="">
            <img src="https://imgapi.twinkle.cat/v0/img/randomcat?5" />
          </GridItem>
          <GridItem colSpan={1} bg="">
            <img src="https://imgapi.twinkle.cat/v0/img/randomcat?15" />
          </GridItem>
        </Grid>
      </Stack>
    </>
  );
}
