import { useState, useEffect } from 'react';
import Pagination from './components/Pagination';
import "./App.css";

export default function App() {
  const ENDPOINT = "http://localhost:4000";
  const limit = 20;
  const [currentPage, setCurrentPage] = useState(1);
  const [totalBooks, setIntProperty] = useState(1);
  useEffect(() => {
    async function fetchData() {
      const response = await fetch(`${ENDPOINT}/catalog/total`);
      const json = await response.json();
      setIntProperty(json.data);
    }
    fetchData();
  }, []);
  const lastPage = Math.floor( totalBooks/limit );;
  
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