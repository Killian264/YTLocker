import React from "react";
import { usePlaylists } from "../shared/api/usePlaylists";
import { VideoListItemController } from "./VideoListItemController";
import { Card } from "../components/Card";
import { useLatestVideos } from "../shared/api/useLatestVideos";
import { LoadingList } from "../components/LoadingList";

export interface VideoListLatestControllerProps {
	className?: string;
	limit: number;
}

export const VideoListLatestController: React.FC<VideoListLatestControllerProps> = ({ className, limit }) => {
	const [loadingP, playlists] = usePlaylists();
	const [loadingV, vidoes] = useLatestVideos();

	if (loadingP || loadingV || playlists == null) {
		return <LoadingList limit={10}></LoadingList>;
	}

	let count = 0;
	let list: JSX.Element[] = [];

	vidoes.forEach((videoId) => {
		playlists.forEach((playlist) => {
			if (!playlist.videos.includes(videoId)) {
				return;
			}

			if (count >= limit) {
				return;
			}
			count++;

			list.push(
				<VideoListItemController
					key={count}
					videoId={videoId}
					playlist={playlist}
				></VideoListItemController>
			);
		});
	});

	return (
		<Card className={className}>
			<div className="flex justify-between -mb-1 -mt-1">
				<div className="text-2xl font-semibold">
					<span className="leading-none -mt-0.5">Videos</span>
				</div>
			</div>
			<div className="grid gap-2">{list}</div>
		</Card>
	);
};
