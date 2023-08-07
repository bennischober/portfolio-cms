"use client";

import { useState } from "react";
import {
    AppShell,
    Header,
    Navbar,
    SegmentedControl,
    Space,
    Text,
    createStyles,
    getStylesRef,
    rem,
} from "@mantine/core";
import {
    IconShoppingCart,
    IconLicense,
    IconMessage2,
    IconBellRinging,
    IconMessages,
    IconFingerprint,
    IconKey,
    IconSettings,
    Icon2fa,
    IconUsers,
    IconFileAnalytics,
    IconDatabaseImport,
    IconReceipt2,
    IconReceiptRefund,
    IconLogout,
    IconSwitchHorizontal,
} from "@tabler/icons-react";
import Link from "next/link";
import { LogoutAction } from "@/actions";

const useStyles = createStyles((theme) => ({
    navbar: {
        backgroundColor:
            theme.colorScheme === "dark" ? theme.colors.dark[7] : theme.white,
    },

    title: {
        textTransform: "uppercase",
        letterSpacing: rem(-0.25),
    },

    link: {
        ...theme.fn.focusStyles(),
        display: "flex",
        alignItems: "center",
        textDecoration: "none",
        fontSize: theme.fontSizes.sm,
        color:
            theme.colorScheme === "dark"
                ? theme.colors.dark[1]
                : theme.colors.gray[7],
        padding: `${theme.spacing.xs} ${theme.spacing.sm}`,
        borderRadius: theme.radius.sm,
        fontWeight: 500,

        "&:hover": {
            backgroundColor:
                theme.colorScheme === "dark"
                    ? theme.colors.dark[6]
                    : theme.colors.gray[0],
            color: theme.colorScheme === "dark" ? theme.white : theme.black,

            [`& .${getStylesRef("icon")}`]: {
                color: theme.colorScheme === "dark" ? theme.white : theme.black,
            },
        },
    },

    linkIcon: {
        ref: getStylesRef("icon"),
        color:
            theme.colorScheme === "dark"
                ? theme.colors.dark[2]
                : theme.colors.gray[6],
        marginRight: theme.spacing.sm,
    },

    linkActive: {
        "&, &:hover": {
            backgroundColor: theme.fn.variant({
                variant: "light",
                color: theme.primaryColor,
            }).background,
            color: theme.fn.variant({
                variant: "light",
                color: theme.primaryColor,
            }).color,
            [`& .${getStylesRef("icon")}`]: {
                color: theme.fn.variant({
                    variant: "light",
                    color: theme.primaryColor,
                }).color,
            },
        },
    },

    footer: {
        borderTop: `${rem(1)} solid ${
            theme.colorScheme === "dark"
                ? theme.colors.dark[4]
                : theme.colors.gray[3]
        }`,
        paddingTop: theme.spacing.md,
    },
}));

const tabs = {
    // thes would be static
    overview: [
        { link: "/schema", label: "Schema", icon: IconBellRinging },
        { link: "/api-endpoints", label: "API Endpoints", icon: IconReceipt2 },
        { link: "/file-bucket", label: "File Bucket", icon: IconFingerprint },
        { link: "/monitoring", label: "Monitoring", icon: IconKey },
        { link: "/test", label: "Databases", icon: IconDatabaseImport },
        { link: "/test", label: "Authentication", icon: Icon2fa },
        { link: "/test", label: "Other Settings", icon: IconSettings },
    ],
    // this would be loaded dynamically
    data: [
        { link: "/test", label: "Orders", icon: IconShoppingCart },
        { link: "/test", label: "Receipts", icon: IconLicense },
        { link: "/test", label: "Reviews", icon: IconMessage2 },
        { link: "/test", label: "Messages", icon: IconMessages },
        { link: "/test", label: "Customers", icon: IconUsers },
        { link: "/test", label: "Refunds", icon: IconReceiptRefund },
        { link: "/test", label: "Files", icon: IconFileAnalytics },
    ],
};

export interface AppNavigationProps {
    children: React.ReactNode;
}

export function AppNavigation({ children }: AppNavigationProps) {
    const { classes, cx } = useStyles();
    const [section, setSection] = useState<"overview" | "data">("overview");
    const [active, setActive] = useState("Home");

    const links = tabs[section].map((item) => (
        <Link
            className={cx(classes.link, {
                [classes.linkActive]: item.label === active,
            })}
            href={item.link}
            key={item.label}
            onClick={(event) => {
                // event.preventDefault();
                setActive(item.label);
            }}
        >
            <item.icon className={classes.linkIcon} stroke={1.5} />
            <span>{item.label}</span>
        </Link>
    ));

    return (
        <AppShell
            padding="md"
            navbar={
                <Navbar width={{ sm: 300 }} p="md" className={classes.navbar}>
                    <Link
                        className={cx(classes.link, {
                            [classes.linkActive]: "Home" === active,
                        })}
                        href="/"
                        onClick={(event) => {
                            // event.preventDefault();
                            setActive("Home");
                        }}
                    >
                        <IconMessage2
                            className={classes.linkIcon}
                            stroke={1.5}
                        />
                        <span>Home</span>
                    </Link>
                    <Space h="xl" />
                    <Navbar.Section>
                        <SegmentedControl
                            value={section}
                            onChange={(value: "overview" | "data") =>
                                setSection(value)
                            }
                            transitionTimingFunction="ease"
                            fullWidth
                            data={[
                                { label: "Overview", value: "overview" },
                                { label: "Data", value: "data" },
                            ]}
                        />
                    </Navbar.Section>

                    <Navbar.Section grow mt="xl">
                        {links}
                    </Navbar.Section>

                    <Navbar.Section className={classes.footer}>
                        <Link
                            href="#"
                            className={classes.link}
                            onClick={(event) => event.preventDefault()}
                        >
                            <IconSwitchHorizontal
                                className={classes.linkIcon}
                                stroke={1.5}
                            />
                            <span>Change account</span>
                        </Link>

                        <Link href="/auth/logout" className={classes.link}>
                            <IconLogout
                                className={classes.linkIcon}
                                stroke={1.5}
                                type="submit"
                            />
                            <span>Logout</span>
                        </Link>
                    </Navbar.Section>
                </Navbar>
            }
        >
            {children}
        </AppShell>
    );
}
