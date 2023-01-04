import PageLink from './PageLink';
import './Pagination.css';
import { getPaginationItems } from '../lib/pagination';
import {TableArea, TableGoodreadsBookProps} from './BookTable';
import { Box, Table, Loader, Center } from '@mantine/core';
import { useForm } from '@mantine/form';
import { SetStateAction, useState } from 'react';
import { TextInput } from '@mantine/core';
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
  var limit = 20;
  var pageNums = getPaginationItems(currentPage, lastPage, maxLength);
  const fetcher = (url: string) =>
  fetch(`${ENDPOINT}/${url}`).then((r) => r.json());
  const [searchValue, setSearchValue] = useState('');
  const [uriValue, setUriValue] = useState({uri: "catalog/books"});
  const {data} = useSWR<TableGoodreadsBookProps>(`${uriValue.uri}?page=${currentPage}&limit=${limit}`, fetcher);

  const handleChange = (e: { target: { value: SetStateAction<string>; }; }) => {
    setSearchValue(e.target.value);
    if(e.target.value===""){
      setUriValue({uri: "catalog/books"});
    } else {
      limit = 5;
      setUriValue({uri: "catalog/search/love"});
    }
  };

  return (
    <nav className="pagination" aria-label="Pagination">
      <TextInput
          placeholder="Search books"
          value={searchValue}
          onChange={handleChange}
      />
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
