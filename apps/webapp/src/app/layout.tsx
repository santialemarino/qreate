import '@qreate/ui/styles';
import '@/app/globals.css';

import type { Metadata } from 'next';

export const metadata: Metadata = {
  title: 'Qreate',
  description: 'QR Code Redirect Service',
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}
