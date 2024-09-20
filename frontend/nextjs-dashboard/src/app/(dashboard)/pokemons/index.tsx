"use client";

import { Button, Card } from "react-bootstrap";
import React from "react";
import { newResource, ResourceCollection } from "@/models/resource";
import { useRouter } from "next/navigation";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faPlus } from "@fortawesome/free-solid-svg-icons";
import useSWR from "swr";
import complaintList from "@/components/Page/complaint/complaintList";
import Cookies from "js-cookie";
import useDictionary from "@/locales/dictionary-hook";

type Props = {
  props: {
    complaintResource: ResourceCollection<complaint>;
    page: number;
    perPage: number;
    sort: string;
    order: string;
  };
};

export default function Index(props: Props) {
  const {
    props: { complaintResource: complaintResourceFallback, page, perPage },
  } = props;

  const router = useRouter();
  const dict = useDictionary();

  const complaintListURL = `http://localhost:8080/get_complain`;
  const url = new URL(complaintListURL);
  url.searchParams.set("offset", page.toString());
  url.searchParams.set("limit", perPage.toString());
  // url.searchParams.set('_sort', sort)
  // url.searchParams.set('_order', order)

  const fetcher = (...args: Parameters<typeof fetch>) =>
    fetch(...args, {
      headers: { Authorization: `Bearer ${Cookies.get("auth")}` },
    }).then(async (res) => {
      if (res.ok) {
        const complaints: any[] = await res.json();
        const total = Number(res.headers.get("x-total-count")) ?? 0;
        return newResource(complaints, total, page, perPage);
      }
      return complaintResourceFallback;
    });

  const { data: complaintResource } = useSWR(url.toString(), fetcher, {
    fallbackData: complaintResourceFallback,
  });

  return (
    <Card>
      <Card.Header>{dict.complaints?.title}</Card.Header>
      <Card.Body>
        <div className="mb-3 text-end">
          <Button
            variant="success"
            onClick={() => router.push("/complaints/create")}
          >
            <FontAwesomeIcon icon={faPlus} fixedWidth />
            {dict.complaints?.add_new}
          </Button>
        </div>
        <complaintList complaints={complaintResource?.data} />
      </Card.Body>
    </Card>
  );
}
