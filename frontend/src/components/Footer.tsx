//ref: https://chakra-templates.dev/page-sections/footer
import {
  Box,
  chakra,
  Container,
  Stack,
  Text,
  useColorModeValue,
  VisuallyHidden,
} from "@chakra-ui/react";
import {
  FaInstagram,
  FaTwitter,
  FaYoutube,
  FaDiscord,
  FaMedium,
  FaGithub,
  FaBook,
} from "react-icons/fa";
import { ReactNode } from "react";

const SocialButton = ({
  children,
  label,
  href,
}: {
  children: ReactNode;
  label: string;
  href: string;
}) => {
  return (
    <chakra.button
      bg={useColorModeValue("blackAlpha.100", "whiteAlpha.100")}
      rounded={"full"}
      w={8}
      h={8}
      cursor={"pointer"}
      as={"a"}
      href={href}
      display={"inline-flex"}
      alignItems={"center"}
      justifyContent={"center"}
      transition={"background 0.3s ease"}
      _hover={{
        bg: useColorModeValue("blackAlpha.200", "whiteAlpha.200"),
      }}
    >
      <VisuallyHidden>{label}</VisuallyHidden>
      {children}
    </chakra.button>
  );
};

export default function SmallWithSocial() {
  return (
    <Box
      bg={useColorModeValue("gray.50", "gray.900")}
      color={useColorModeValue("gray.700", "gray.200")}
    >
      <Container
        as={Stack}
        maxW={"6xl"}
        py={4}
        direction={{ base: "column", md: "row" }}
        spacing={4}
        justify={{ base: "center", md: "space-between" }}
        align={{ base: "center", md: "center" }}
      >
        <Text>© 2022 Twinkle. All rights reserved</Text>
        <Stack direction={"row"} spacing={6}>
          <SocialButton label={"Twitter"} href={"/twitter"}>
            <FaTwitter />
          </SocialButton>
          <SocialButton label={"Discord"} href={"/discord"}>
            <FaDiscord />
          </SocialButton>
          <SocialButton label={"Medium"} href={"#"}>
            <FaMedium />
          </SocialButton>
          <SocialButton label={"Github"} href={"#"}>
            <FaGithub />
          </SocialButton>
          <SocialButton label={"Gitbook"} href={"/docs"}>
            <FaBook />
          </SocialButton>
          {/* <SocialButton label={"YouTube"} href={"#"}>
            <FaYoutube />
          </SocialButton>
          <SocialButton label={"Instagram"} href={"#"}>
            <FaInstagram />
          </SocialButton> */}
        </Stack>
      </Container>
    </Box>
  );
}
