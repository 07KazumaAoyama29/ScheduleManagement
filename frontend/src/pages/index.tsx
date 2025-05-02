import useSWR from "swr";

const fetcher = (url: string) => fetch(url).then((r) => r.json());

export default function Home() {
  const { data } = useSWR(`${process.env.NEXT_PUBLIC_API_URL}/events`, fetcher);

  if (!data) return <p>Loading...</p>;

  return (
    <ul className="p-4 space-y-1">
      {data.map((ev: any) => (
        <li key={ev.id} className="border p-2 rounded">
          {ev.title} — {new Date(ev.start_at).toLocaleString()}
        </li>
      ))}
    </ul>
  );
}