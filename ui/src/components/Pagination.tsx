import PageLink from './PageLink';
import './Pagination.css';
import { getPaginationItems } from '../lib/pagination';
import {TableArea, TableGoodreadsBookProps} from './BookTable';
import { Box, Table, Loader, Center } from '@mantine/core';
import { useForm } from '@mantine/form';
import useSWR from "swr";

export type Props = {
  currentPage: number;
  lastPage: number;
  maxLength: number;
  setCurrentPage: (page: number) => void;
};

export default function Pagination({
  currentPage,
  lastPage,
  maxLength,
  setCurrentPage,
}: Props) {

  const ENDPOINT = "http://localhost:4000";
  const limit = 20;
  const pageNums = getPaginationItems(currentPage, lastPage, maxLength);
  const fetcher = (url: string) =>
  fetch(`${ENDPOINT}/${url}`).then((r) => r.json());
  const { data } = useSWR<TableGoodreadsBookProps>(`catalog/books?page=${currentPage}&limit=${limit}`, fetcher);

  return (
    <nav className="pagination" aria-label="Pagination">
      <Box>
       {data ? <TableArea data={data.data} /> : <Loader />}
      </Box>
      <PageLink
        disabled={currentPage === 1}
        onClick={() => setCurrentPage(currentPage - 1)}
      >
        Previous
      </PageLink>
      {pageNums.map((pageNum, idx) => (
        <PageLink
          key={idx}
          active={currentPage === pageNum}
          disabled={isNaN(pageNum)}
          onClick={() => setCurrentPage(pageNum)}
        >
          {!isNaN(pageNum) ? pageNum : '...'}
        </PageLink>
      ))}
      <PageLink
        disabled={currentPage === lastPage}
        onClick={() => setCurrentPage(currentPage + 1)}
      >
        Next
      </PageLink>
    </nav>
  );
}