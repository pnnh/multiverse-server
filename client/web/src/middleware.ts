import {NextResponse} from 'next/server'

export function middleware (request: Request) {
  const requestHeaders = new Headers(request.headers)
  requestHeaders.set('x-url', request.url)

  //console.log('requestHeaders', request.url)

  return NextResponse.next({
    request: {
      headers: requestHeaders,
    }
  })
}

export const config = {
  matcher: '/((?!api|_next/static|_next/image|favicon.ico).*)',
}
