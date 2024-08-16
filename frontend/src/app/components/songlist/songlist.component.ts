import { Component, OnInit } from '@angular/core';
import { SongListService } from '../../services/songlist.service';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-songlist',
  templateUrl: './songlist.component.html',
  styleUrls: ['./songlist.component.css']
})
export class SongListComponent implements OnInit {
  songLists: any[] = [];
  songListForm: FormGroup;

  constructor(private songListService: SongListService, private fb: FormBuilder) {
    this.songListForm = this.fb.group({
      name: ['', Validators.required],
      songs: this.fb.array([]),
    });
  }

  ngOnInit(): void {
    this.loadSongLists();
  }

  loadSongLists(): void {
    this.songListService.getSongLists().subscribe(data => {
      this.songLists = data;
    });
  }

  createSongList(): void {
    const newSongList = this.songListForm.value;
    this.songListService.createSongList(newSongList).subscribe(() => {
      this.loadSongLists();
    });
  }

  deleteSongList(id: number): void {
    this.songListService.deleteSongList(id).subscribe(() => {
      this.loadSongLists();
    });
  }
}
