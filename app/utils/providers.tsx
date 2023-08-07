"use client";

import { AuthProvider } from "@/hooks/useAuth";
import { MantineProvider } from "@mantine/core";
import { Notifications } from "@mantine/notifications";

export function Providers({ children }: { children: React.ReactNode }) {
    return (
        <MantineProvider
            theme={{ colorScheme: "dark" }}
            withGlobalStyles
            withNormalizeCSS
        >
            <AuthProvider>
                <Notifications />
                {children}
            </AuthProvider>
        </MantineProvider>
    );
}
