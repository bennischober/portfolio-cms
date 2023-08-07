"use client";

import { useAuth } from "@/hooks/useAuth";
import {
    TextInput,
    PasswordInput,
    Paper,
    Title,
    Text,
    Container,
    Button,
} from "@mantine/core";
import { useForm, zodResolver } from "@mantine/form";
import { useRouter } from "next/navigation";
import { z } from "zod";

export interface User {
    username: string;
    password: string;
}

interface AuthenticationLoginProps {
    login: (user: User) => Promise<any>;
}

const schema = z.object({
    username: z.string().email({ message: "Invalid email format" }),
});

export function AuthenticationLogin({ login }: AuthenticationLoginProps) {
    const Form = useForm<User>({
        initialValues: {
            username: "",
            password: "",
        },

        validate: zodResolver(schema),
    });

    const { setToken, token } = useAuth();
    const router = useRouter();

    return (
        <form
            onSubmit={Form.onSubmit(async (user) => {
                const data = await login(user);
                if (data.token) {
                    setToken(data.token);
                    router.push("/");
                }
            })}
        >
            <Container size={420} my={40}>
                <Title
                    align="center"
                    sx={(theme) => ({
                        fontFamily: `Greycliff CF, ${theme.fontFamily}`,
                        fontWeight: 900,
                    })}
                >
                    Welcome back!
                </Title>
                <Text color="dimmed" size="sm" align="center" mt={5}>
                    Please login to continue!
                </Text>

                <Paper withBorder shadow="md" p={30} mt={30} radius="md">
                    <TextInput
                        label="Email"
                        placeholder="you@mantine.dev"
                        required
                        {...Form.getInputProps("username")}
                    />
                    <PasswordInput
                        label="Password"
                        placeholder="Your password"
                        required
                        mt="md"
                        {...Form.getInputProps("password")}
                    />
                    <Button type="submit" fullWidth mt="xl">
                        Sign in
                    </Button>
                </Paper>
            </Container>
        </form>
    );
}
