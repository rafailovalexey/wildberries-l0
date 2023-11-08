import '@/app/globals.css'

import type { Metadata } from 'next'

export const metadata: Metadata = {
  title: 'Wildberries L0',
  description: 'Wildberries L0',
}

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
      <html className="h-auto" lang="en">
          <body className="h-auto">{children}</body>
      </html>
  );
}

