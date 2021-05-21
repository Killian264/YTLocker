import { useQuery } from "react-query";
import { useHistory } from "react-router";
import { DROPLET_BASE } from "../env";
import { useBearer } from "../hooks/useBearer";
import { Video } from "../types";
import { ApiResponse2 } from "./api";

export interface VideoFetchResponse extends ApiResponse2 {
	video: Video | null;
}

export const useVideo = (id: number, enabled = true): [boolean, Video | null] => {
	const [bearer] = useBearer("");
	let history = useHistory();

	const { isLoading, isError, data } = useQuery(["channel", id], () => fetchVideo(bearer, id), {
		refetchInterval: false,
		enabled: enabled,
	});

	if (isLoading || isError || data === undefined) {
		return [true, null];
	}

	if (data.status === 401) {
		history.push("/login");
	}

	if (data.video === null) {
		return [true, null];
	}

	return [false, data.video];
};

const fetchVideo = async (bearer: string, id: number): Promise<VideoFetchResponse> => {
	const url = DROPLET_BASE + `/video/${id}`;

	const options = {
		method: "GET",
		headers: {
			Authorization: bearer,
		},
	};

	const res = await (await fetch(url, options)).json();

	let video: VideoFetchResponse = {
		status: res.Status,
		message: res.Message,
		video: null,
	};

	if (video.status === 200) {
		video.video = parseVideo(res.Data);
	}

	return video;
};

const parseVideo = (data: any): Video => {
	let thumbnail = data.Thumbnails[data.Thumbnails.length - 1];

	let parsed: Video = {
		id: data.ID,
		youtubeId: data.YoutubeID,
		thumbnailUrl: thumbnail.URL,
		title: data.Title,
		description: data.Description,
		created: new Date(Date.parse(data.CreatedAt)),
	};

	return parsed;
};
