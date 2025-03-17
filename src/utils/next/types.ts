import { ReactNode } from "react";

/**
 * https://nextjs.org/docs/app/api-reference/file-conventions/page#props
 */
export type PageProps<
  TParams extends Record<string, string | Array<string>> = Record<string, never>,
> = {
  params: Promise<TParams>;
  searchParams: Promise<Record<string, string | Array<string>>>;
};

/**
 * https://nextjs.org/docs/app/api-reference/file-conventions/layout#props
 */
export type LayoutProps<
  TParams extends Record<string, string | Array<string>> = Record<string, never>,
> = {
  children: ReactNode;
  params: Promise<TParams>;
};

/**
 * https://nextjs.org/docs/app/api-reference/file-conventions/error#props
 */
export type ErrorPageProps = {
  error: Error & { digest?: string };
  reset: () => void;
};
