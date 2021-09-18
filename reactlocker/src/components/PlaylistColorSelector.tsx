import React, { useState } from "react";
import { Color, Playlist } from "../shared/types";
import { ColorSelector } from "./ColorSelector";

export interface PlaylistColorSelectorProps {
	className?: string;
	OnClick: (color: Color) => void;
	playlists: Playlist[];
	selectedPlaylist?: Playlist | null;
}

export const PlaylistColorSelector: React.FC<PlaylistColorSelectorProps> = ({
	className,
	OnClick,
	playlists,
	selectedPlaylist = null,
}) => {
	const [color, setColor] = useState<Color | null>(
		selectedPlaylist === null ? null : selectedPlaylist.color
	);

	const disabledColors = playlists
		.filter((p) => {
			return selectedPlaylist == null || p.id !== selectedPlaylist.id;
		})
		.map((playlist) => {
			return playlist.color;
		});

	return (
		<ColorSelector
			className={className}
			onClick={(color) => {
				setColor(color);
				OnClick(color);
			}}
			disabled={disabledColors}
			selected={color}
		></ColorSelector>
	);
};
