import '@qreate/ui/styles';
import '@/app/globals.css';

import type { Metadata } from 'next';

export const metadata: Metadata = {
  title: 'Qreate Backoffice',
  description: 'QR Code Management Platform',
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}
