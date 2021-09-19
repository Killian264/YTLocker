import { useState } from "react";
import { usePlaylistList } from "../hooks/api/usePlaylistList";
import { PlaylistCreateCard } from "../components/PlaylistCreateCard";
import { PlaylistView } from "../components/PlaylistView";
import { usePlaylistDelete } from "../hooks/api/usePlaylistDelete";
import { usePlaylistUpdate } from "../hooks/api/usePlaylistUpdate";
import { usePlaylist } from "../hooks/api/usePlaylist";

export interface PlaylistViewControllerProps {
	className?: string;
	playlistId: number;
	BackClick: () => void;
}

export const PlaylistViewController: React.FC<PlaylistViewControllerProps> = ({
	className,
	playlistId,
	BackClick,
}) => {
	const [isLoadingPlaylists, playlists] = usePlaylistList();
	const [isLoadingPlaylist, playlist] = usePlaylist(playlistId);

	const [isEditing, setEditing] = useState(false);

	const deletePlaylist = usePlaylistDelete();
	const updatePlaylist = usePlaylistUpdate();

	if (isLoadingPlaylist || isLoadingPlaylists || playlist == null) {
		return <div>Loading...</div>;
	}

	const swap = () => {
		setEditing(!isEditing);
	};

	if (isEditing) {
		return (
			<PlaylistCreateCard
				className={className}
				editPlaylist={playlist}
				playlists={playlists}
				CreateClick={(title, description, color) => {
					updatePlaylist(playlist.id, title, description, color);
				}}
				BackClick={swap}
			></PlaylistCreateCard>
		);
	}

	return (
		<PlaylistView
			className={className}
			playlist={playlist}
			EditClick={swap}
			DeleteClick={() => {
				deletePlaylist(playlistId);
				BackClick();
			}}
			BackClick={BackClick}
		></PlaylistView>
	);
};
