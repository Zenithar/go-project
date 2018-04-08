package todo

import "testing"

func TestGetProjectReq_Validate(t *testing.T) {
	type fields struct {
		ID string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Empty",
			fields: fields{
				ID: "",
			},
			wantErr: true,
		},
		{
			name: "Invalid length",
			fields: fields{
				ID: "123456789",
			},
			wantErr: true,
		},
		{
			name: "Invalid charset (utf-8)",
			fields: fields{
				ID: "JKéMWXCVBNAZERTYUIOPQSDFGHJKLMWX",
			},
			wantErr: true,
		},
		{
			name: "Invalid charset (symbol)",
			fields: fields{
				ID: "JK-MWXCVBNAZERTYUIOPQSDFGHJKLMWX",
			},
			wantErr: true,
		},
		{
			name: "Valid",
			fields: fields{
				ID: "JKAMWXCVBNAZERTYUIOPQSDFGHJKLMWX",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &GetProjectReq{
				Id: tt.fields.ID,
			}
			if err := m.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("GetProjectReq.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
