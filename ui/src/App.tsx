
// import { Box, Table, Loader, Center } from '@mantine/core';
// import { useForm } from '@mantine/form';
// import useSWR from "swr";
import { useState } from 'react';
import Pagination from './components/Pagination';
import "./App.css";

export const ENDPOINT = "http://localhost:4000";

const fetcher = (url: string) =>
  fetch(`${ENDPOINT}/${url}`).then((r) => r.json());

export interface TableGoodreadsBookProps {
	data: {
    book_id: number;
    title: string;
    authors: string;
    average_rating: number;
    isbn: string;
    isbn13: string;
    language_code: string;
    num_pages: string;
    ratings_count: number;
    text_reviews_count: number;
    publication_date: string;
    publisher: string;
	}[];
}

export function TableArea({ data }: TableGoodreadsBookProps) {

	const rows = data.map(row => (
		<tr key={row.book_id}>
      <td>{row.title}</td>
      <td>{row.authors}</td>
      <td>{row.isbn}</td>
      <td>{row.average_rating}</td>
      <td>{row.num_pages}</td>
		</tr>
	));

	return (
		<Table horizontalSpacing="sm" verticalSpacing="sm">
				<thead>
					<tr>
          <th>Title</th>
          <th>Authors</th>
          <th>ISBN</th>
          <th>Average Rating</th>
          <th>Number of Pages</th>
					</tr>
				</thead>
        
				<tbody>{rows}</tbody>
		</Table>
	);
}


// function App() {
//   const { data } = useSWR<TableGoodreadsBookProps>("catalog/books", fetcher);
//   return (
//     <Box>
//       {data ? <TableArea data={data.data} /> : <Loader />}
//     </Box>
//   )
// }
// export default App;


export default function App() {
  const [currentPage, setCurrentPage] = useState(1);
  const lastPage = 20;

  return (
    <div className="container">
      <h1>Catalog of Books</h1>
      <Pagination
        currentPage={currentPage}
        lastPage={lastPage}
        maxLength={7}
        setCurrentPage={setCurrentPage}
      />
    </div>
  );
}


