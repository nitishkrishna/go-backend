import { Box, Table, Loader, Center } from '@mantine/core';

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