import { NextResponse } from 'next/server';

export function middleware(request: Request) {
    // this adds the route path to the request headers => can be accessed in RootLayout
    const requestHeaders = new Headers(request.headers);
    requestHeaders.set('x-url', request.url);

    return NextResponse.next({
        request: {
            headers: requestHeaders,
        }
    });
}