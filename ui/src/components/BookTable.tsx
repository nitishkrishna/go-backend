import { Box, Table, Loader, Center } from '@mantine/core';

export interface TableGoodreadsBookProps {
     data: {
     title: string;
     authors: string;
     average_rating: number;
     num_pages: string;
     }[];
}
 
export function TableArea({ data }: TableGoodreadsBookProps) {
 
     const rows = data.map(row => (
         <tr key={row.title}>
       <td>{row.title}</td>
       <td>{row.authors}</td>
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
           <th>Average Rating</th>
           <th>Number of Pages</th>
                     </tr>
                 </thead>
         
                 <tbody>{rows}</tbody>
         </Table>
     );
}
