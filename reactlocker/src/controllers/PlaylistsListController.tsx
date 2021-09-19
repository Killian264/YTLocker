import React from "react";
import { Card } from "../components/Card";
import { PlaylistListItem } from "../components/PlaylistListItem";
import { BuildPlaylistUrl } from "../shared/urls";
import { usePlaylistList } from "../hooks/api/usePlaylistList";
import { LoadingList } from "../components/LoadingList";
import { Plus, SvgBox } from "../components/Svg";

export interface PlaylistsListControllerProps {
	className?: string;
	PlaylistClick: (id: number) => void;
	CreatePlaylistClick: () => void;
}

export const PlaylistsListController: React.FC<PlaylistsListControllerProps> = ({
	className,
	CreatePlaylistClick,
	PlaylistClick,
}) => {
	const [isLoadingPlaylists, playlists] = usePlaylistList();

	if (isLoadingPlaylists) {
		return <LoadingList limit={3}></LoadingList>;
	}

	const list = playlists.map((playlist, index) => {
		return (
			<PlaylistListItem
				url={BuildPlaylistUrl(playlist.youtubeId)}
				key={index}
				playlist={playlist}
				onClick={() => {
					PlaylistClick(playlist.id);
				}}
			></PlaylistListItem>
		);
	});

	return (
		<Card className={className}>
			<div className="flex justify-between -mb-2 -mt-2">
				<div className="text-2xl font-semibold">
					<span className="leading-none -mt-0.5">Playlists</span>
				</div>
				<SvgBox className={`border-green-500 p-0.5`} onClick={CreatePlaylistClick}>
					<Plus className="text-green-500" size={24}></Plus>
				</SvgBox>
			</div>
			<div className="grid gap-2">{list}</div>
		</Card>
	);
};
