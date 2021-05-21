import React from "react";
import { Card } from "../components/Card";
import { PlaylistListItem } from "../components/PlaylistListItem";
import { BuildPlaylistUrl } from "../shared/urls";
import { PlusButton } from "../components/PlusButton";
import { usePlaylists } from "../shared/api/usePlaylists";
import { LoadingList } from "../components/LoadingList";

export interface PlaylistsListControllerProps {
	className?: string;
}

export const PlaylistsListController: React.FC<PlaylistsListControllerProps> = ({ className }) => {
	const [loading, playlists] = usePlaylists();

	if (loading) {
		return <LoadingList limit={3}></LoadingList>;
	}

	const list = playlists.map((playlist, index) => {
		return (
			<PlaylistListItem
				url={BuildPlaylistUrl(playlist.youtubeId)}
				key={index}
				playlist={playlist}
			></PlaylistListItem>
		);
	});

	return (
		<Card className={className}>
			<div className="flex justify-between -mb-1 -mt-1">
				<div className="text-2xl font-semibold">
					<span className="leading-none -mt-0.5">Playlists</span>
				</div>
				<PlusButton color="secondary" disabled={false}></PlusButton>
			</div>
			<div className="grid gap-2">{list}</div>
		</Card>
	);
};
